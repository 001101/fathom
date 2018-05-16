'use strict';

import Pikaday from 'pikaday';
import { h, Component } from 'preact';

class Pikadayer extends Component {
  componentDidMount() {
    new Pikaday({ 
      field: this.base,
      onSelect: this.props.onSelect,
      position: 'bottom right',
   })
  }
  render(props) {
    return <input {...props} />
  }
}

export default Pikadayer
