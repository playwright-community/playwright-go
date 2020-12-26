package playwright

func createObjectFactory(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) interface{} {
	switch objectType {
	case "Android":
		return nil
	case "BindingCall":
		return newBindingCall(parent, objectType, guid, initializer)
	case "Browser":
		return newBrowser(parent, objectType, guid, initializer)
	case "BrowserType":
		return newBrowserType(parent, objectType, guid, initializer)
	case "BrowserContext":
		return newBrowserContext(parent, objectType, guid, initializer)
	case "ConsoleMessage":
		return newConsoleMessage(parent, objectType, guid, initializer)
	case "Dialog":
		return newDialog(parent, objectType, guid, initializer)
	case "Download":
		return newDownload(parent, objectType, guid, initializer)
	case "ElementHandle":
		return newElementHandle(parent, objectType, guid, initializer)
	case "Frame":
		return newFrame(parent, objectType, guid, initializer)
	case "JSHandle":
		return newJSHandle(parent, objectType, guid, initializer)
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
	default:
		panic(objectType)
	}
}
