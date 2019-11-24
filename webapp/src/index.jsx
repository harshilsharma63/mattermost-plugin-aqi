import React from 'react';
import Actions from './actions';
import Reducers from './reducers';
import Util from './utils';

import LeftSidebarHeader from "./components/left_sidebar_header";

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
                const pollutionData = {
                    id: Util.uuidv4(),
                    data: JSON.parse(event.data.pollutionData),
                };
                store.dispatch(Actions.receivePollutionData(pollutionData));
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
