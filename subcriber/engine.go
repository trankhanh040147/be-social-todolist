package subscriber

import (
	"context"
	"log"
	"social-todo-list/common"
	"social-todo-list/common/asyncjob"
	"social-todo-list/pubsub"

	goservice "github.com/200Lab-Education/go-sdk"
)

type subJob struct {
	Name string
	Hdl  func(ctx context.Context, msg *pubsub.Message) error
}

type GroupJob interface {
	Run(ctx context.Context) error
}

type pbEngine struct {
	serviceCtx goservice.ServiceContext
}

func NewPBEngine(serviceCtx goservice.ServiceContext) *pbEngine {
	return &pbEngine{serviceCtx: serviceCtx}
}

func (engine *pbEngine) Start() error {
	engine.startSubTopic(common.TopicUserLikedItem, true,
		IncreaseLikedCountAfterUserLikeItem(engine.serviceCtx),
		PushNotiAfterUserLikeItem(engine.serviceCtx),
	)

	engine.startSubTopic(common.TopicUserUnlikedItem, true,
		DecreaseLikedCountAfterUserUnlikeItem(engine.serviceCtx),
	)

	return nil
}

func (engine *pbEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, jobs ...subJob) error {
	ps := engine.serviceCtx.MustGet(common.PluginPubSub).(pubsub.PubSub)

	c, _ := ps.Subscribe(context.Background(), topic)
	for _, item := range jobs {
		log.Println("Setup subscriber for", item.Name)
	}

	getJobHandler := func(job *subJob, msg *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Printf("Running job[%s] - Value: %v", job.Name, msg.Data())
			return job.Hdl(ctx, msg)
		}
	}

	go func() {
		for {
			msg := <-c

			jobHdls := make([]asyncjob.Job, len(jobs))
			for index := range jobs {
				jobHdlIdnex := getJobHandler(&jobs[index], msg)
				jobHdls[index] = asyncjob.NewJob(jobHdlIdnex, asyncjob.WithName(jobs[index].Name))
			}

			group := asyncjob.NewGroup(isConcurrent, jobHdls...)
			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}
