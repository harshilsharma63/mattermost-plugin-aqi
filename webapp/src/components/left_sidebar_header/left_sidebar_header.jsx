import React from 'react';
import PropTypes from 'prop-types';
import Client from '../../client';
import Constants from "../../constants";


export default class LeftSidebarHeader extends React.Component {

    constructor(props) {
        super(props);
        this.state = LeftSidebarHeader.getInitialState();
    }

    static getInitialState() {
        return {
            pollutionData: {
                id: "",
                data: [],
            }
        };
    }

    componentDidMount() {
        Client.doGet(Constants.URL_REFRESH_DATA).then(r => {
            console.log(r);
        });
    }

    componentDidUpdate(prevProps, prevState, snapshot) {
        if (this.props.pollutionData.id !== prevProps.pollutionData.id) {
            this.setState({
                pollutionData: this.props.pollutionData,
            })
        }
    }

    getPollutionRating = (aqi) => {
        if (aqi <= 50) {
            return "Good";
        } else if (aqi <= 100) {
            return "Moderate";
        } else if (aqi <= 150) {
            return "Unhealthy for Sensitive Groups";
        } else if (aqi <= 200) {
            return "Unhealthy";
        } else if (aqi <= 300) {
            return "Very Unhealthy";
        } else if (aqi <= 500) {
            return "Hazardous";
        } else {
            return "Beyond Scale"
        }
    };

    getPollutionColor = (aqi) => {
        if (aqi <= 50) {
            return "#a8e05f";
        } else if (aqi <= 100) {
            return "#fdd74b";
        } else if (aqi <= 150) {
            return "#fe9b57";
        } else if (aqi <= 200) {
            return "#fe6a69";
        } else if (aqi <= 300) {
            return "#a97abc";
        } else if (aqi <= 500) {
            return "#a87383";
        } else {
            return "#ff0000"
        }
    };

    getPrimaryPollutantDisplayName = (pollutantCode) => {
        const pollutantDisplayNames = {
            "p2": "PM2.5",
            "p1": "PM10",
            "o3": "OZONE",
            "n2": "NO2",
            "s2": "SO2",
            "co": "CO",
        };

        return pollutantDisplayNames[pollutantCode];
    };

    render() {
        let rows = [];

        if (this.state.pollutionData.data === undefined) {
            return null;
        }

        for (let i=0; i < this.state.pollutionData.data.length; ++i) {
            let cityData = this.state.pollutionData.data[i];
            let aqiColorCode = this.getPollutionColor(cityData.data.current.pollution.aqius);
            let style = getStyle(this.props.theme, aqiColorCode);

            rows.push(
                <div key={i} className={'city'} style={style.city}>
                    <div className={'header'} style={style.header}>
                        <div className={'city_name'}>{cityData.data.city}</div>
                        <div className={'city_weather'} style={style.cityWeather}>
                            <img
                                className={'city_weather_icon'}
                                src={`https://www.airvisual.com/images/${cityData.data.current.weather.ic}.png`}
                                alt={' '}
                                style={style.cityWeatherIcon}
                            />
                            <span className={'city_temperature'}>
                                {cityData.data.current.weather.tp} °C
                            </span>
                        </div>
                    </div>
                    <div className={'footer'} style={style.footer}>
                        <div className={'city_aqi'} style={style.cityAqi}>
                            <span className={'city_aqi_value'}
                                  style={style.cityAqiValue}>{cityData.data.current.pollution.aqius}</span>
                            <span className={'city_aqi_type'} style={style.cityAqiType}>{'US AQI'}</span>
                        </div>
                        <div className={'pollution_rating'} style={style.rating}>
                            <div
                                className={'city_pollution_rating'}
                                style={style.cityPollutionRating}>{this.getPollutionRating(cityData.data.current.pollution.aqius)}</div>
                            <div
                                className={'primary_pollutant'}
                                style={style.primaryPollutant}>{this.getPrimaryPollutantDisplayName(cityData.data.current.pollution.mainus)}</div>
                        </div>
                    </div>
                </div>
            );
        }

        return (
            <div>
                {rows}
            </div>
        );
    }
}

LeftSidebarHeader.propTypes = {
    pollutionData: PropTypes.object.isRequired,
    theme: PropTypes.object,
};

const getStyle = (theme, aqiColorCode) => {
    return {
        city: {
            margin: "4px",
            border: "1px solid gray",
            boxShadow: "0px 0px 2px #666",
            borderRadius: "5px"
        },
        header: {
            display: "flex",
            padding: "5px 10px",
            backgroundColor: "black",
            color: "white",
            borderRadius: "5px 5px 0 0",
        },
        footer: {
            display: "flex",
            flexDirection: "row",
            height: "unset",
            backgroundColor: aqiColorCode,
            color: "#0000007a",
            borderRadius: "0 0 5px 5px",
        },
        cityWeather: {
            marginLeft: "auto",
        },
        cityWeatherIcon: {
            width: "20px",
            marginRight: "5px",
        },
        cityAqi: {
            width: "80px",
            display: "flex",
            flexDirection: "column",
            textAlign: "center",
        },
        cityAqiValue: {
            fontSize: "200%",
        },
        cityAqiType: {
            fontSize: "70%",
        },
        rating: {
            display: "flex",
            flexDirection: "column",
            flexGrow: "1",
            textAlign: "center",
            padding: "0 10px",
            justifyContent: "center",
        },
        cityPollutionRating: {
            fontSize: "100%",
        },
        primaryPollutant: {
            fontSize: "70%",
        }

    };
};
