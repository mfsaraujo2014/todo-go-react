import React from 'react';
import {CgClose, CgPen} from 'react-icons/cg'
import './Task.css'
import { useHistory } from 'react-router-dom';

const Task = ({task, handleTaskClick, handleTaskDeletion}) => {
    const history = useHistory();

    const handleTaskDetailsClick = () => {
        history.push(`/${task.title}`)
    }
    return (
        <div 
        className="task-container" 
        style={task.completed ? {borderLeft:'30px solid greenyellow'} : {}}
        >
            <div className="task-title" onClick={() => handleTaskClick(task.id)}>
                {task.title}
            </div>

            <div className="buttons-container">
                <button className="remove-task-button" onClick={() => handleTaskDeletion(task.id)}>
                    <CgClose/>
                </button>
                <button className="edit-task-button" onClick={handleTaskDetailsClick} >
                    <CgPen/>
                </button>
            </div>
        </div>
    )
}
 
export default Task;