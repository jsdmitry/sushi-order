import React, {Component} from 'react';
import './App.css';
import MainLayout from './components/mainLayout';

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
