package worker

import "crawler/engine"

type CrawlService struct {
}

func (CrawlService) Process(req engine.Request, result *engine.ParseResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}

	engineRes, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}

	*result = SerializeParseResult(engineRes)
	return nil
}
