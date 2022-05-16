import React, { useState, useEffect } from 'react';
import axios from 'axios';

export default function Response(props) {
    const [ username, setName ] = useState();
    const [ age, setAge ] = useState();
    const [ character, setCharacter ] = useState({
      CharacterName: '',
      Strength: '',
      Dexterity: '',
      Constitution: '',
      Intelligence: '',
      Wisdom: '',
      Charisma: ''
    });

    
  useEffect(() => {
    const getResponse = async () => {
        const { data } = await axios("/users");
        let name = data.username;
        let age = data.age;
        setName(name);
        setAge(age);
    };
    getResponse();
  }, []);

  useEffect(() => {
    const getResponse = async (endpoint) => {
        const { data } = await axios(endpoint);
        let char = await data;
        setCharacter({
          CharacterName: char.name,
          Strength: char.str,
          Dexterity: char.dex,
          Constitution: char.con,
          Intelligence: char.int,
          Wisdom: char.wis,
          Charisma: char.cha
        });
    };
    getResponse("char");
  }, []);

  return <div style={{textAlign: 'left'}}>
      Name: { username } <br/>
      Age: { age } <br/> <br/> <br/>
      Character Name: {character.CharacterName} <br/>
      Strength: {character.Strength} <br/>
      Dexterity: {character.Dexterity} <br/>
      Constitution: {character.Constitution} <br/>
      Intelligence: {character.Intelligence} <br/>
      Wisdom: {character.Wisdom} <br/>
      Charisma: {character.Charisma} <br/>
    </div>;
}
