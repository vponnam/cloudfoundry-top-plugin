// Copyright (c) 2016 ECS Team, Inc. - All Rights Reserved
// https://github.com/ECSTeam/cloudfoundry-top-plugin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eventDetailView

import "github.com/ecsteam/cloudfoundry-top-plugin/ui/uiCommon/views/helpView"

const HelpText = HelpOverviewText +
	helpView.HelpHeaderText +
	HelpColumnsText +
	helpView.HelpChildLevelDataViewKeybindings +
	helpView.HelpCommonDataViewKeybindings

const HelpOverviewText = `
**Event Details Stats View**

Event detail list view shows a list of event details of the selected
event type and origin.  
`
const HelpColumnsText = `
**Event list stats:**

  DNAME - BOSH deployment name
  JOB_NAME - BOSH job name
  JOB_IDX - BOSH job index
  IP - IP address where event originated
  COUNT - Number of events that have occured
`
