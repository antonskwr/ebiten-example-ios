//
//  ViewController.m
//  EbitenTest
//
//  Created by Keeper on 08/06/2020.
//  Copyright © 2020 Keeper. All rights reserved.
//

#import "ViewController.h"
#import <Foundation/Foundation.h>
#import <Mygame/Mygame.h>

@implementation ViewController {}

- (void)onErrorOnGameUpdate:(NSError *)err {
    // You can define your own error handling e.g., using Crashlytics.
    NSLog(@"Inovation Error!: %@", err);
}

- (void)viewDidLoad {
    [super viewDidLoad];
    
    MygameInitGame((long)self.view.frame.size.width, (long)self.view.frame.size.height);
    
}
@end
