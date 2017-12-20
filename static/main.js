const elements = {
  app: $('#app'),
  messageForm: $('#message-form'),
  userForm: $('#user-form'),
}

const roomId = elements.app.attr('room-id')
const defaultUserId = elements.app.attr('user-id')

$('#chat-form').submit((e) => {
  const { messageForm, userForm } = elements

  if (!userForm.val()) {
    userForm.val(defaultUserId)
  }

  e.preventDefault()
  $.post(`/room/${roomId}`, {
    message: messageForm.val(),
    user_id: userForm.val()
  })

  messageForm.val('')
  messageForm.focus()
})

if (!!window.EventSource) {
  const source = new EventSource(`/stream/${roomId}`)
  source.addEventListener('message', e => {
    $('#messages').append(`${e.data} <br/>`)
    $('html, body').animate({ scrollTop: $(document).height() }, 'slow')
  }, false);
} else {
    alert('Not supported!')
}

$('#message-form').focus()
