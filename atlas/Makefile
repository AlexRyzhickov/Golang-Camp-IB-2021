.PHONY: gen-services
gen-services:
	@atlas init-app -name=portal -helm -kind -delve -gateway
	@atlas init-app -name=responder -helm -kind -delve -pubsub
	@atlas init-app -name=storage -helm -kind -delve  -pubsub -db