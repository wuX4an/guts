import u from '/static/plugins/umbrella.js'



u('body').addClass('bg-green-200')

u('#container').addClass('flex flex-col items-center gap-x-4 p-4')
u('#image').addClass('w-2/12 pt-14')
u('#message').addClass('text-zinc-700 text-2xl font-bold pt-12');

// floro
u('#floro').addClass('text-zinc-600 italic font-medium pb-3 pt-2')
u('#floro-f').addClass('text-rose-500 font-bold')
u('#floro-b').addClass('text-violet-500 font-bold')

//server
u('#server').addClass('text-violet-600 pt-3 font-medium font-sans')

//client
u('#client').addClass('text-rose-600 pt-3 font-medium font-sans')
const url = "https://randomuser.me/api"
const response = await fetch(url);
const json = await response.json();
const title = json.results[0].name.title
const firstName = json.results[0].name.first
const secondName = json.results[0].name.last
u('#client').text(`${title} ${firstName} ${secondName}`)
