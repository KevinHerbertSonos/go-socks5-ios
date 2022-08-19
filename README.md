   Initial wrapper to run a SOCKS5 proxy on iOS.
    
    The main directory contains a simple test application to
	try on your desktop.
    
    Build the iOS directory using gomobile.
    
    https://pkg.go.dev/golang.org/x/mobile/cmd/gomobile
    
    gomobile init
    gomobule bind -target ios
    
    Will produce a package called Socks.xcframework which
	you can include in your iOS project.
