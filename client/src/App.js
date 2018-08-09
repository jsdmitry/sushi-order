import React, {Component} from 'react';
import './App.css';
import MainLayout from './containers/mainLayout';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = { inputType: "string" };
  }

  render() {
    return (
      <MainLayout/>
    );
  }
}

export default App;
