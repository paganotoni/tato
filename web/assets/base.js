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

function actionHTML(action) {
    return `<li class="bg-white border-l-2 border-red-300 p-2 shadow-sm flex flex-row">
        <span class="flex-grow">${action.full}</span> 
        <span class="delete underline cursor-pointer" data-id="${action.ID}">Delete</span>
    </li>`
}

function deleteAction(id) {
    fetch('/api/actions/destroy/'+id, {
        method: 'DELETE',
    }).then(response => {
        console.log(response.status)
    })

    refreshActions();
}

function refreshActions() {
    let actionsList = document.querySelector('#actions-list');
    actionsList.innerHTML = '';
    
    fetch('/api/actions/list').then(response => {
        return response.json();
    }).then(data => {
        document.getElementById("actions-count").innerHTML = data.length

        data.forEach(action => {
            actionsList.insertAdjacentHTML('beforeend', actionHTML(action));
        });

        document.querySelectorAll('#actions-list li .delete').forEach(span => {
            span.addEventListener('click', (e) => deleteAction(e.target.dataset.id));
        })
    })

    fetch('/api/stats/k1-distribution').then(response => {
        return response.json();
    }).then(data => {
        let total =  0
        for (const key in data) {
            total += data[key].total
        }

        document.getElementById("total-attacks").innerHTML = total;

        for (const key in data) {
            data[key]["percentage"] = data[key].total / total;
            data[key]["%="] = data[key].error / data[key].total;
            data[key]["%*"] = data[key].perfect / data[key].total;
            
            let totel = document.getElementById(`zone-${key}-total`)
            if(totel != null) {
                totel.innerHTML = data[key].total;
            }

            let perel = document.getElementById(`zone-${key}-%`)
            if(perel != null) {
                perel.innerHTML = roundToTwo(data[key]["percentage"] * 100) + '%';
            }

            let errel = document.getElementById(`zone-${key}-=%`)
            if(errel != null) {
                errel.innerHTML = roundToTwo(data[key]["%="] * 100) + '%';
            }

            let pefel = document.getElementById(`zone-${key}-*%`)
            if(pefel != null) {
                pefel.innerHTML = roundToTwo(data[key]["%*"] * 100) + '%';
            }
        }
    })

    
}

function roundToTwo(num) {
    return +(Math.round(num + "e+2")  + "e-2");
}
