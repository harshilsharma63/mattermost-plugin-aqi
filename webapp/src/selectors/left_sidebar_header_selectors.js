import Constants from '../constants';

const getPluginState = (state) => state[`plugins-${Constants.PLUGIN_NAME}`] || {};

export const getPollutionData = (state) => getPluginState(state).pollutionData || {};
