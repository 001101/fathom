'use strict';

import { h, Component } from 'preact';
import Client from '../lib/client.js';
import { bind } from 'decko';

class Realtime extends Component {

  constructor(props) {
    super(props)

    this.state = {
      count: 0
    }
  }

  componentDidMount() {
      this.fetchData();
      this.interval = window.setInterval(this.fetchData, 15000);
  }

  componentWillUnmount() {
      clearInterval(this.interval);
  }

  @bind
  fetchData() {
    Client.request(`visitors/count/realtime`)
      .then((d) => { this.setState({ count: d })})
  }

  render(props, state) {
    let visitorText = state.count == 1 ? 'visitor' : 'visitors';
    return (
        <span><span class="count">{state.count}</span> <span>current {visitorText}</span></span>
    )
  }
}

export default Realtime
