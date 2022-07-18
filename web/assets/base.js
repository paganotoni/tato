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
    return `<li class="bg-white border-l-2 border-red-300 p-2 shadow-sm flex flex-row gap-1 items-center">
        <span class="text-xs rounded-full bg-red-100 p-0.5 px-1">P6</span>
        <span class="text-xs rounded-full bg-blue-100 p-0.5 px-1">P3</span>
        <span class="flex-grow">${action.full}</span> 
        <a class="delete cursor-pointer text-red-700" data-id="${action.ID}">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 pointer-events-none" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
            </svg>
        </a>
    </li>`
}

function deleteAction(id) {
    fetch('/api/actions/destroy/'+id, {
        method: 'DELETE',
    }).then(response => {
        if(response.status != 200){
            return
        }

        refreshActions();
    })
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
            span.addEventListener('click', (e) => {
                deleteAction(span.dataset.id)
            });
        })
    })

    fetch('/api/stats/k1').then(response => {
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
                perel.innerHTML = roundTo(1, data[key]["percentage"] * 100) + '%';
            }

            let errel = document.getElementById(`zone-${key}-=%`)
            if(errel != null) {
                errel.innerHTML = roundTo(1, data[key]["%="] * 100) + '%';
            }

            let pefel = document.getElementById(`zone-${key}-*%`)
            if(pefel != null) {
                pefel.innerHTML = roundTo(1, data[key]["%*"] * 100) + '%';
            }
        }
    })

    
}

function roundTo(dec, num) {
    return +(Math.round(num + "e+"+dec)  + "e-"+dec);
}
