document.addEventListener('DOMContentLoaded', () => {
    refreshActions();

    let cmdInput = document.querySelector('#cmd-input');
    cmdInput.addEventListener('keydown', (e) => {
        if (e.keyCode != 13) {
            return
        }
        
        let value = cmdInput.value;
        cmdInput.value = '';

        fetch('/api/actions/create',{
            method: 'POST',
            body: JSON.stringify({
                input: value
            })
        }).then(response => {
            if(response.status != 200){
                return
            }

            refreshActions();
        })
    })
});

function refreshActions() {
    let actionsList = document.querySelector('#actions-list');
    actionsList.innerHTML = '';
    
    fetch('/api/actions/list').then(response => {
        return response.json();
    }).then(data => {
        document.getElementById("actions-count").innerHTML = data.length
        
        data.forEach(action => {
            actionsList.insertAdjacentHTML('beforeend', `<li class="bg-white border-l-2  border-red-300 p-2 shadow-sm">${action.full}</li>`);
        });
    })
}