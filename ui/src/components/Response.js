import React, { useState, useEffect } from 'react';
import axios from 'axios';

function Response(props) {
    const [username, setName] = useState();
    const [age, setAge] = useState();

    const getResponse = async (user = "users") => {
        let url  = "http://localhost:8080/"+user;
        const { data } = await axios(url);
        let name = data.username;
        let age = data.age;
        setName(name);
        setAge(age);
    };
    
  useEffect(() => {
    getResponse();
  }, []);

  return <div>
      Name: { username } <br/>
      Age: { age }
    </div>;
}

export default Response