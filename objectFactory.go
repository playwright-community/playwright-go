package playwright

func createObjectFactory(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) interface{} {
	switch objectType {
	case "Android":
		return nil
	case "Artifact":
		return newArtifact(parent, objectType, guid, initializer)
	case "BindingCall":
		return newBindingCall(parent, objectType, guid, initializer)
	case "Browser":
		return newBrowser(parent, objectType, guid, initializer)
	case "BrowserType":
		return newBrowserType(parent, objectType, guid, initializer)
	case "BrowserContext":
		return newBrowserContext(parent, objectType, guid, initializer)
	case "CDPSession":
		return newCDPSession(parent, objectType, guid, initializer)
	case "ConsoleMessage":
		return newConsoleMessage(parent, objectType, guid, initializer)
	case "Dialog":
		return newDialog(parent, objectType, guid, initializer)
	case "ElementHandle":
		return newElementHandle(parent, objectType, guid, initializer)
	case "Frame":
		return newFrame(parent, objectType, guid, initializer)
	case "JSHandle":
		return newJSHandle(parent, objectType, guid, initializer)
	case "LocalUtils":
		return nil
	case "Tracing":
		return newTracing(parent, objectType, guid, initializer)
	case "APIRequestContext":
		return nil
	case "Page":
		return newPage(parent, objectType, guid, initializer)
	case "Playwright":
		return newPlaywright(parent, objectType, guid, initializer)
	case "Request":
		return newRequest(parent, objectType, guid, initializer)
	case "Response":
		return newResponse(parent, objectType, guid, initializer)
	case "Route":
		return newRoute(parent, objectType, guid, initializer)
	case "WebSocket":
		return newWebsocket(parent, objectType, guid, initializer)
	case "Worker":
		return newWorker(parent, objectType, guid, initializer)
	case "Selectors":
		return nil
	case "Electron":
		return nil
	case "FetchRequest":
		c := &channelOwner{}
		c.createChannelOwner(c, parent, objectType, guid, initializer)
		return c
	case "JsonPipe":
		return newJsonPipe(parent, objectType, guid, initializer)
	case "Stream":
		return newStream(parent, objectType, guid, initializer)
	default:
		panic(objectType)
	}
}
