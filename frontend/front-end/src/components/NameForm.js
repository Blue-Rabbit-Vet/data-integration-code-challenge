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
                <div>
                    <input type="text" name='firstName' placeholder="First Name"    value={this.state.firstName} onChange={this.handleChange}/>
                </div>
                <div>
                    <input type="text" name='lastName'  placeholder="Last Name"     value={this.state.lastName} onChange={this.handleChange} />
                </div>
                <div>
                    <input type="text" name='number'    placeHolder="Number"        value={this.state.number} onChange={this.handleChange}/>
                </div>
                <input type="submit" value="Submit" />
            </form>
        );
    }
}

export default NameForm
