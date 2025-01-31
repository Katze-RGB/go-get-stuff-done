import streamlit as st
import requests,os,base64
rainbowcat = open('rainbowcat.gif','rb')
contents = rainbowcat.read()
data_url = base64.b64encode(contents).decode("utf-8")
rainbowcat.close()

st.title('Go Get Stuff Done Magic Scheduler app whee')
st.markdown(
    f'<img src="data:image/gif;base64,{data_url}" alt="meow meow meow">',
    unsafe_allow_html=True,
)
if st.button('Whats Up Next???'):
    data = requests.get(url='http://backend:3000/get_next_task').json()
    print(data)
    prio = data['eng_priority']
    id = data['id']
    desc = data['description']
    est_time = data['estimated_length']
    try:
        st.write('ID: '+str(id))
        st.write('Priority: '+prio)
        st.write('Description: '+desc)
        st.write('Estimated Time: '+str(est_time)+" mins")
    except KeyError:
        print(data)
        st.write(data['detail'])

st.title('Create a new task')
desc = st.text_input('task description')
prio = st.radio('priority',['high','medium','low'])
est = st.number_input('estimated time (in whole mins)', min_value=1, max_value=60, value=5)

if st.button('Create Task!'):
    prio_out=0
    if prio == 'high':
        prio_out = 3
    elif prio == 'medium':
        prio_out = 2
    elif prio == 'low':
        prio_out = 1

    data = {"description":desc, "estimated_length":est, "priority":prio_out}
    headers = {'Content-type':'application/json'}
    requests.post("http://backend:3000/todo_task/", json=data, headers=headers)

st.title('Complete or Delete a task')
task_id = st.number_input('Task ID', min_value=1)

if st.button('Complete Task!'):
    data = requests.post("http://backend:3000/complete_task/"+str(task_id))
    try:
        st.write((data))
    except KeyError:
        print(data)
        st.write(data['detail'])

if st.button('Delete Task :<'):
    data = requests.delete("http://backend:3000/todo_task/"+str(task_id))
    try:
        st.write((data))
    except KeyError:
        print(data)
        st.write(data['detail'])

st.title("Get Status Report")
status_date = st.date_input('date for productivity report', format='YYYY-MM-DD')
if st.button('Get Status Report'):
    data = requests.get("http://backend:3000/productivity_report/"+str(status_date)).json()
    try:
        st.write('date: '+str(data['date']))
        st.write('Tasks Completed: '+str(data['tasks_completed']))
        st.write('Mins Spent: '+str(data['mins_spent']))
    except KeyError:
        print(data)
        st.write(data['detail'])
