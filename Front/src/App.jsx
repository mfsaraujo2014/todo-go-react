import React, { useEffect, useState } from "react";
import Tasks from "./components/Tasks";
import './App.css';
import AddTask from "./components/AddTask";
//import {v4 as uuidv4} from 'uuid';
import Header from './components/Header'
import {BrowserRouter as Router, Route} from 'react-router-dom'
//import api from "./services/api";
import axios from "axios";

const App = () => {
  const [tasks, setTasks] = useState([
    {
      id: 1,
      title: 'Estudar programaçao',
      completed: false,
    },
  ]);

  useEffect(() => {
    const fetchTasks = async () => {
      const {data} = await axios.get("http://localhost:5000/tasks");
      setTasks(data);
    }

    fetchTasks()
  }, [])

  const handleTaskClick = (taskId) => {
    const newTasks = tasks.map((task) => {
      if (task.id === taskId) return {...task, completed: !task.completed}

      return task;
    })

    setTasks(newTasks)
  }

  const handleTaskAddition = (taskTitle) => {
    /*const newTasks = [
      ...tasks,
      {
        title: taskTitle,
        id: uuidv4(),
        completed: false,
      },
    ];*/
    const res = axios.post("http://localhost:5000/tasks", {
    title: taskTitle,
    id: 1,
    completed: false})
    setTasks(res);
    console.log(res)
  };

  

  

  const handleTaskDeletion = (taskId) => {
    axios.delete(`http://localhost:5000/tasks/${taskId}`)
    //const newTasks = tasks.filter(task => task.id !== taskId)
    setTasks(tasks.filter(task => task.id !== taskId))
  };

  return (
    <Router>
      <div className="container">
        <Header />
          <Route path="/" exact render={() => (
            <>
              <AddTask handleTaskAddition={handleTaskAddition} />
              <Tasks tasks={tasks} handleTaskClick={handleTaskClick} handleTaskDeletion={handleTaskDeletion} />
            </>
          )} 
          />
          <Route path="/:taskTitle"/>
      </div>
    </Router>
    
  )
}

export default App;