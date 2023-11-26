package playwright

func createObjectFactory(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) interface{} {
	switch objectType {
	case "Android":
		return nil
	case "AndroidSocket":
		return nil
	case "AndroidDevice":
		return nil
	case "APIRequestContext":
		return newAPIRequestContext(parent, objectType, guid, initializer)
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
	case "Dialog":
		return newDialog(parent, objectType, guid, initializer)
	case "Electron":
		return nil
	case "ElectronApplication":
		return nil
	case "ElementHandle":
		return newElementHandle(parent, objectType, guid, initializer)
	case "Frame":
		return newFrame(parent, objectType, guid, initializer)
	case "JSHandle":
		return newJSHandle(parent, objectType, guid, initializer)
	case "JsonPipe":
		return newJsonPipe(parent, objectType, guid, initializer)
	case "LocalUtils":
		localUtils := newLocalUtils(parent, objectType, guid, initializer)
		if localUtils.connection.localUtils == nil {
			localUtils.connection.localUtils = localUtils
		}
		return localUtils
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
	case "Selectors":
		return newSelectorsOwner(parent, objectType, guid, initializer)
	case "SocksSupport":
		return nil
	case "Stream":
		return newStream(parent, objectType, guid, initializer)
	case "Tracing":
		return newTracing(parent, objectType, guid, initializer)
	case "WebSocket":
		return newWebsocket(parent, objectType, guid, initializer)
	case "Worker":
		return newWorker(parent, objectType, guid, initializer)
	case "WritableStream":
		return newWritableStream(parent, objectType, guid, initializer)
	default:
		panic(objectType)
	}
}
