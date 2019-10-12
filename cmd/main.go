package main

import(
"github.com/soxueren/greenplum-node/pkg/routers"
"github.com/soxueren/greenplum-node/pkg/server"
"github.com/soxueren/greenplum-node/pkg/setting"
"github.com/DeanThompson/ginpprof"
)

func main() {
	setting.Setup()

	srv := server.Start()
	router := routers.InitRouter()
	ginpprof.Wrap(router)
	server.RegistryRouter(srv, router)
}