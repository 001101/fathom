'use strict'

import { h, Component } from 'preact';
import LogoutButton from '../components/LogoutButton.js';
import Realtime from '../components/Realtime.js';
import DatePicker from '../components/DatePicker.js';
import CountWidget from '../components/CountWidget.js';
import Table from '../components/Table.js';

import { bind } from 'decko';

class Dashboard extends Component {
  constructor(props) {
    super(props)

    this.state = {
      period: (window.location.hash.substring(2) || 'week'),
      before: 0,
      after: 0,
    }
  }

  @bind
  changePeriod(s) {
    this.setState({ period: s.period, before: s.before, after: s.after })
    window.history.replaceState(this.state, null, `#!${s.period}`)
  }

  render(props, state) {
    return (
    <div class="wrapper">

      <header class="section">
        <nav class="main-nav animated fadeInDown">
            <ul>
              <li class="logo"><a href="/">Fathom</a></li>
              <li class="visitors"><Realtime /></li>
              <li class="spacer">&middot;</li>
              <li class="signout"><LogoutButton onSuccess={this.props.onLogout} /></li>
          </ul>
        </nav>
      </header>

      <section class="section animated fadeInUp delayed_02s">
        <nav class="date-nav">
          <DatePicker onChange={this.changePeriod} value={state.period} />
        </nav>

        <div class="boxes">
          <div class="box box-totals animated fadeInUp delayed_03s">
            <CountWidget title="Unique visitors" endpoint="stats/site/visitors" before={state.before} after={state.after} />
            <CountWidget title="Page views" endpoint="stats/site/pageviews" before={state.before} after={state.after} />
            <CountWidget title="Avg time on site" endpoint="stats/site/duration" format="duration" before={state.before} after={state.after} />
            <CountWidget title="Bounce rate" endpoint="stats/site/bounces" format="percentage" before={state.before} after={state.after} />
          </div>
  
          <Table endpoint="stats/pages" headers={["Top pages", "Views", "Uniques"]} before={state.before} after={state.after} />
          <Table endpoint="stats/referrers" headers={["Top referrers", "Views", "Uniques"]} before={state.before} after={state.after} showHostname="true" />

        </div>
      </section>

      <footer class="section"></footer>
    </div>
  )}
}

export default Dashboard
