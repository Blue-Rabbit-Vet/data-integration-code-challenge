import React from 'react';
import axios from 'axios';

class NameForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {firstName: '', lastName: '', number: ''};

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange(event) {
        console.log(event)
        this.setState({
            [event.target.name] : event.target.value
        })
    }
    handleSubmit(event) {
        event.preventDefault();
        axios.post('http://localhost:8080/addplayer', {
            firstName: this.state.firstName,
            lastName: this.state.lastName,
            number: this.state.number
        })
            .then((response) => {
                console.log(response);
            }, (error) => {
                console.log(error);
            });
    }

    render() {
        return (
            <form onSubmit={this.handleSubmit}>
                <input type="text" name='firstName' value={this.state.firstName} onChange={this.handleChange}/>
                <input type="text" name='lastName' value={this.state.lastName} onChange={this.handleChange} />
                <input type="text" name='number' value={this.state.number} onChange={this.handleChange}/>
                <input type="submit" value="Submit" />
            </form>
        );
    }
}

export default NameForm
