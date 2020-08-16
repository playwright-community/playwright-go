package playwright

func createObjectFactory(parent *ChannelOwner, objectType string, guid string, initializer interface{}) interface{} {
	switch objectType {
	case "Playwright":
		return newPlaywright(parent, objectType, guid, initializer)
	case "BrowserType":
		return newBrowserType(parent, objectType, guid, initializer)
	case "Browser":
		return newBrowser(parent, objectType, guid, initializer)
	case "BrowserContext":
		return newBrowserContext(parent, objectType, guid, initializer)
	case "Frame":
		return newFrame(parent, objectType, guid, initializer)
	case "Page":
		return newPage(parent, objectType, guid, initializer)
	case "Request":
		return newRequest(parent, objectType, guid, initializer)
	case "Response":
		return newResponse(parent, objectType, guid, initializer)
	case "Selectors":
		return nil
	case "Electron":
		return nil
	default:
		panic(objectType)
	}
}
