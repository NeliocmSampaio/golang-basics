package router

func (r *Router) SetupRouter() {
	r.engine.GET("/health", r.healthController.HandleHealth)
}
