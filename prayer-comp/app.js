import React from 'react';
import './PrayerTimeApp.css';

class PrayerTimeApp extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      // set initial prayer times
      fajr: '5:00am',
      dhuhr: '12:00pm',
      asr: '3:00pm',
      maghrib: '5:00pm',
      isha: '7:00pm',
    };
  }

  // fetch prayer times from api
  componentDidMount() {
    fetch('/api/prayer-times')
    .then(res => res.json())
    .then(data => {
      // set new prayer times
      this.setState({
        fajr: data.fajr,
        dhuhr: data.dhuhr,
        asr: data.asr,
        maghrib: data.maghrib,
        isha: data.isha
      });
    });
  }

  render() {
    return (
      <div className="prayer-time-app">
        <h2>Prayer Times</h2>
        <ul>
          <li>Fajr: {this.state.fajr}</li>
          <li>Dhuhr: {this.state.dhuhr}</li>
          <li>Asr: {this.state.asr}</li>
          <li>Maghrib: {this.state.maghrib}</li>
          <li>Isha: {this.state.isha}</li>
        </ul>
      </div>
    );
  }
}

export default PrayerTimeApp;