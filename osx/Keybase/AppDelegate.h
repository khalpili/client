//
//  AppDelegate.h
//  Keybase
//
//  Created by Gabriel on 12/11/14.
//  Copyright (c) 2014 Gabriel Handford. All rights reserved.
//

#import <Cocoa/Cocoa.h>

#import "KBRPClient.h"
#import <KBKeybase/KBKeybase.h>
#import "KBRPC.h"
#import "KBAppView.h"
#import <GHKit/GHKit.h>

@interface AppDelegate : NSObject <NSApplicationDelegate>

@property KBAPIClient *APIClient;

+ (KBAppView *)appView;

+ (AppDelegate *)sharedDelegate;

- (IBAction)preferences:(id)sender;
- (IBAction)quit:(id)sender;

+ (void)setError:(NSError *)error sender:(NSView *)sender;
- (void)setFatalError:(NSError *)error;

+ (NSString *)bundleFile:(NSString *)file;

+ (NSString *)applicationSupport:(NSArray *)subdirs create:(BOOL)create error:(NSError **)error;

- (void)openURLString:(NSString *)URLString sender:(NSView *)sender;

@end
