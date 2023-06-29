package product

import (
	"context"
)

type SyncProduct struct {
}

func (s *SyncProduct) RunSkillProduct(ctx context.Context) error {
	// rabbitMqQueue := consts.SkillProductQueues
	// skillProduct, err := rabbitmq.ConsumeMessage(ctx, rabbitMqQueue)
	// if err != nil {
	// 	return err
	// }
	// var forever chan struct{}
	//
	// go func() {
	// 	for d := range skillProduct {
	// 		log.LogrusObj.Infof("Received run skill like : %s", d.Body)
	//
	// 		// 落库
	// 		reqRabbitMQ := new(types.SkillProductReq)
	// 		err = json.Unmarshal(d.Body, reqRabbitMQ)
	// 		if err != nil {
	// 			log.LogrusObj.Infof("Received run story like : %s", err)
	// 		}
	//
	// 		err = service.SkillProductMQ2MySQL(ctx, reqRabbitMQ)
	// 		if err != nil {
	// 			log.LogrusObj.Infof("Received run story like : %s", err)
	// 		}
	//
	// 	}
	// }()
	//
	// log.LogrusObj.Infoln(err)
	// <-forever

	return nil
}
