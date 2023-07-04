package config

type AppConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `ini:"etcd"`
}

type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Time    int    `ini:"time"`
}

//type TailLogConf struct {
//	Filename string `ini:"filename"`
//}
