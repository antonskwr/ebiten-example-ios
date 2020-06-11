//
//  AppDelegate.m
//  EbitenTest
//
//  Created by Keeper on 08/06/2020.
//  Copyright © 2020 Keeper. All rights reserved.
//

#import "AppDelegate.h"
#import "ViewController.h"

@interface AppDelegate ()

@end

@implementation AppDelegate


- (BOOL)application:(UIApplication *)application didFinishLaunchingWithOptions:(NSDictionary *)launchOptions {
    CGRect screenRect = [[UIScreen mainScreen] bounds];
    CGFloat screenWidth = screenRect.size.width;
    CGFloat screenHeight = screenRect.size.height;
    
    MobileSetScreenSize((long)screenWidth, (long)screenHeight);
    
    self.window = UIWindow.new;
    ViewController *vc = ViewController.new;
    
    self.window.rootViewController = vc;
    [self.window makeKeyAndVisible];
    return YES;
}

@end
