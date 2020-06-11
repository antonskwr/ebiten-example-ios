//
//  ViewController.m
//  EbitenTest
//
//  Created by Keeper on 08/06/2020.
//  Copyright © 2020 Keeper. All rights reserved.
//

#import "ViewController.h"
#import <Foundation/Foundation.h>

@implementation ViewController {}

- (void)onErrorOnGameUpdate:(NSError *)err {
    // You can define your own error handling e.g., using Crashlytics.
    NSLog(@"Game Error!: %@", err);
}

- (void)viewDidLoad {
    [super viewDidLoad];
//    MobileSetScreenSize((long)self.view.frame.size.width, (long)self.view.frame.size.height);
}

@end
