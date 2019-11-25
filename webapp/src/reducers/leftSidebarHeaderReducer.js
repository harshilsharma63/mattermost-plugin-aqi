import Constants from '../constants';

let prevPollutionData;

export const pollutionData = (state = false, action) => {
    switch (action.type) {
    case Constants.ACTIONS.RECEIVE_POLLUTION_DATA:
        prevPollutionData = action.pollutionData;
        return action.pollutionData;
    default:
        return prevPollutionData || {};
    }
};
