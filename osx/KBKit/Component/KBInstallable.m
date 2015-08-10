//
//  KBInstallable.m
//  Keybase
//
//  Created by Gabriel on 5/18/15.
//  Copyright (c) 2015 Gabriel Handford. All rights reserved.
//

#import "KBInstallable.h"

@interface KBInstallableComponent ()
@property KBEnvConfig *config;
@end

@implementation KBInstallableComponent

- (instancetype)initWithConfig:(KBEnvConfig *)config {
  if ((self = [super init])) {
    _config = config;
  }
  return self;
}

- (void)setComponentStatus:(KBComponentStatus *)componentStatus {
  _componentStatus = componentStatus;
  [self componentDidUpdate];
}

- (void)componentDidUpdate { }

- (GHODictionary *)componentStatusInfo {
  if (!_componentStatus) return [GHODictionary dictionary];
  GHODictionary *info = [GHODictionary dictionary];

  info[@"Status Error"] = _componentStatus.error.localizedDescription;
  info[@"Install Status"] = NSStringFromKBInstallStatus(_componentStatus.installStatus);
  info[@"Runtime Status"] = NSStringFromKBRuntimeStatus(_componentStatus.runtimeStatus);

  return info;
}

@end
