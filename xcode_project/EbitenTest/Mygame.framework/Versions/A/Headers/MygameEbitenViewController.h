// Code generated by ebitenmobile. DO NOT EDIT.

#import <UIKit/UIKit.h>

@interface MygameEbitenViewController : UIViewController

// onErrorOnGameUpdate is called on the main thread when an error happens when updating a game.
// You can define your own error handler, e.g., using Crashlytics, by overwriting this method.
- (void)onErrorOnGameUpdate:(NSError*)err;

// suspendGame suspends the game.
// It is recommended to call this when the application is being suspended e.g.,
// UIApplicationDelegate's applicationWillResignActive is called.
- (void)suspendGame;

// resumeGame resumes the game.
// It is recommended to call this when the application is being resumed e.g.,
// UIApplicationDelegate's applicationDidBecomeActive is called.
- (void)resumeGame;

@end
