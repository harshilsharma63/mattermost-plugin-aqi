import React from 'react';

import LeftSidebarHeader from "./components/left_sidebar_header/left_sidebar_header";

import Constants from './constants';

//
// Define the plugin class that will register
// our plugin components.
//
class PluginClass {
    initialize(registry) {
        registry.registerLeftSidebarHeaderComponent(LeftSidebarHeader);
    }
}

//
// To register your plugin you must expose it
// on window.
//
window.registerPlugin(Constants.PLUGIN_NAME, new PluginClass());
