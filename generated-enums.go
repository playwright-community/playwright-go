package playwright

func getMixedState(in string) *MixedState {
	v := MixedState(in)
	return &v
}

type MixedState string

var (
	MixedStateOn    *MixedState = getMixedState("On")
	MixedStateOff               = getMixedState("Off")
	MixedStateMixed             = getMixedState("Mixed")
)

func getColorScheme(in string) *ColorScheme {
	v := ColorScheme(in)
	return &v
}

type ColorScheme string

var (
	ColorSchemeLight        *ColorScheme = getColorScheme("light")
	ColorSchemeDark                      = getColorScheme("dark")
	ColorSchemeNoPreference              = getColorScheme("no-preference")
)

func getForcedColors(in string) *ForcedColors {
	v := ForcedColors(in)
	return &v
}

type ForcedColors string

var (
	ForcedColorsActive *ForcedColors = getForcedColors("active")
	ForcedColorsNone                 = getForcedColors("none")
)

func getReducedMotion(in string) *ReducedMotion {
	v := ReducedMotion(in)
	return &v
}

type ReducedMotion string

var (
	ReducedMotionReduce       *ReducedMotion = getReducedMotion("reduce")
	ReducedMotionNoPreference                = getReducedMotion("no-preference")
)

func getServiceWorkerPolicy(in string) *ServiceWorkerPolicy {
	v := ServiceWorkerPolicy(in)
	return &v
}

type ServiceWorkerPolicy string

var (
	ServiceWorkerPolicyAllow *ServiceWorkerPolicy = getServiceWorkerPolicy("allow")
	ServiceWorkerPolicyBlock                      = getServiceWorkerPolicy("block")
)

func getHarNotFound(in string) *HarNotFound {
	v := HarNotFound(in)
	return &v
}

type HarNotFound string

var (
	HarNotFoundAbort    *HarNotFound = getHarNotFound("abort")
	HarNotFoundFallback              = getHarNotFound("fallback")
)

func getMouseButton(in string) *MouseButton {
	v := MouseButton(in)
	return &v
}

type MouseButton string

var (
	MouseButtonLeft   *MouseButton = getMouseButton("left")
	MouseButtonRight               = getMouseButton("right")
	MouseButtonMiddle              = getMouseButton("middle")
)

func getKeyboardModifier(in string) *KeyboardModifier {
	v := KeyboardModifier(in)
	return &v
}

type KeyboardModifier string

var (
	KeyboardModifierAlt     *KeyboardModifier = getKeyboardModifier("Alt")
	KeyboardModifierControl                   = getKeyboardModifier("Control")
	KeyboardModifierMeta                      = getKeyboardModifier("Meta")
	KeyboardModifierShift                     = getKeyboardModifier("Shift")
)

func getScreenshotAnimations(in string) *ScreenshotAnimations {
	v := ScreenshotAnimations(in)
	return &v
}

type ScreenshotAnimations string

var (
	ScreenshotAnimationsDisabled *ScreenshotAnimations = getScreenshotAnimations("disabled")
	ScreenshotAnimationsAllow                          = getScreenshotAnimations("allow")
)

func getScreenshotCaret(in string) *ScreenshotCaret {
	v := ScreenshotCaret(in)
	return &v
}

type ScreenshotCaret string

var (
	ScreenshotCaretHide    *ScreenshotCaret = getScreenshotCaret("hide")
	ScreenshotCaretInitial                  = getScreenshotCaret("initial")
)

func getScreenshotScale(in string) *ScreenshotScale {
	v := ScreenshotScale(in)
	return &v
}

type ScreenshotScale string

var (
	ScreenshotScaleCss    *ScreenshotScale = getScreenshotScale("css")
	ScreenshotScaleDevice                  = getScreenshotScale("device")
)

func getScreenshotType(in string) *ScreenshotType {
	v := ScreenshotType(in)
	return &v
}

type ScreenshotType string

var (
	ScreenshotTypePng  *ScreenshotType = getScreenshotType("png")
	ScreenshotTypeJpeg                 = getScreenshotType("jpeg")
)

func getElementState(in string) *ElementState {
	v := ElementState(in)
	return &v
}

type ElementState string

var (
	ElementStateVisible  *ElementState = getElementState("visible")
	ElementStateHidden                 = getElementState("hidden")
	ElementStateStable                 = getElementState("stable")
	ElementStateEnabled                = getElementState("enabled")
	ElementStateDisabled               = getElementState("disabled")
	ElementStateEditable               = getElementState("editable")
)

func getWaitForSelectorState(in string) *WaitForSelectorState {
	v := WaitForSelectorState(in)
	return &v
}

type WaitForSelectorState string

var (
	WaitForSelectorStateAttached *WaitForSelectorState = getWaitForSelectorState("attached")
	WaitForSelectorStateDetached                       = getWaitForSelectorState("detached")
	WaitForSelectorStateVisible                        = getWaitForSelectorState("visible")
	WaitForSelectorStateHidden                         = getWaitForSelectorState("hidden")
)

func getWaitUntilState(in string) *WaitUntilState {
	v := WaitUntilState(in)
	return &v
}

type WaitUntilState string

var (
	WaitUntilStateLoad             *WaitUntilState = getWaitUntilState("load")
	WaitUntilStateDomcontentloaded                 = getWaitUntilState("domcontentloaded")
	WaitUntilStateNetworkidle                      = getWaitUntilState("networkidle")
	WaitUntilStateCommit                           = getWaitUntilState("commit")
)

func getLoadState(in string) *LoadState {
	v := LoadState(in)
	return &v
}

type LoadState string

var (
	LoadStateLoad             *LoadState = getLoadState("load")
	LoadStateDomcontentloaded            = getLoadState("domcontentloaded")
	LoadStateNetworkidle                 = getLoadState("networkidle")
)

func getMedia(in string) *Media {
	v := Media(in)
	return &v
}

type Media string

var (
	MediaScreen *Media = getMedia("screen")
	MediaPrint         = getMedia("print")
	MediaNull          = getMedia("null")
)

func getSameSiteAttribute(in string) *SameSiteAttribute {
	v := SameSiteAttribute(in)
	return &v
}

type SameSiteAttribute string

var (
	SameSiteAttributeStrict *SameSiteAttribute = getSameSiteAttribute("Strict")
	SameSiteAttributeLax                       = getSameSiteAttribute("Lax")
	SameSiteAttributeNone                      = getSameSiteAttribute("None")
)
