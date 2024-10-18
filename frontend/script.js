$(document).ready(function () {
    console.log('is ready')

    $('.btn').on('click', function () {
        let target = $(this).data('target')

        $('.video-wrapper div').hide()
        $('.video-wrapper video').trigger('pause')
        $(target).show()

        $(target + ' video').trigger('play')
    })
})