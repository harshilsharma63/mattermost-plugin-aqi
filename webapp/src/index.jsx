import React from 'react';

import Actions from './actions';
import Reducers from './reducers';

import LeftSidebarHeader from './components/left_sidebar_header';

import Constants from './constants';

//
// Define the plugin class that will register
// our plugin components.
//
class PluginClass {
    initialize(registry, store) {
        registry.registerLeftSidebarHeaderComponent(LeftSidebarHeader);

        registry.registerWebSocketEventHandler(
            `custom_${Constants.PLUGIN_NAME}_receive_pollution_data`,
            (event) => {
                console.log(event.data);
                store.dispatch(Actions.receivePollutionData(JSON.parse(event.data.data)));
            },
        );

        registry.registerReducer(Reducers);
    }
}

//
// To register your plugin you must expose it
// on window.
//
window.registerPlugin(Constants.PLUGIN_NAME, new PluginClass());
