function clickHandler(event) {
    // selection
    // click word

    // selection
    let selection = window.getSelection().toString();
    if (selection.length > 0) {
        getDictResult(selection, displayMeaning);
        return;
    }

    // click word
    let range;
    let textNode;
    let offset;
    if (document.caretRangeFromPoint) {
        range = document.caretRangeFromPoint(event.clientX, event.clientY);
        textNode = range.startContainer;
        offset = range.startOffset;
    } else if (document.caretPositionFromPoint) {
        let pos = document.caretPositionFromPoint(event.clientX, event.clientY);
        textNode = pos.offsetNode;
        offset = pos.offset;
    }

    if (offset >= 0 && offset < textNode.length && textNode.textContent[offset].match(/[a-zA-Z]/) != null) {
        // find word
        var text = textNode.textContent;
        var start = text.substring(0, offset).search(/\w+$/);
        var end = text.substring(offset).search(/\W/);
        if (end < 0) {
            end = text.length;
        }
        var word = text.substring(start, end + offset);
        if (word.length > 0) {
            getDictResult(word, displayMeaning);
        }
        return
    }

    // blank
    blankHandler(event);
}

function blankHandler(event) {
    console.log("blank");
}

function getDictResult(word, callback) {
    const regex = /[\w']+/g;
    word = word.match(regex);
    if (word == null) {
        return;
    }
    word = word[0];

    var url = 'https://api.dictionaryapi.dev/api/v2/entries/en/' + word;
    var xhr = new XMLHttpRequest();
    xhr.open('GET', url, true);
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4 && xhr.status == 200) {
            var json = JSON.parse(xhr.responseText);
            if (json.length == 0) {
                return;
            }
            callback(json)
        }
    }
    xhr.send();
}

function displayMeaning(json) {
    if (json == null) {
        return;
    }
    json = json[0];
    const word = json['word'];
    // choose the phonetic with audio
    let phonetic_map = {}
    const phonetics = json['phonetics'];
    for (let i = 0; i < phonetics.length; i++) {
        if (!phonetics[i]['audio']) {
            continue;
        }
        phonetic = phonetics[i]['text'];
        audio = phonetics[i]['audio'];
        // if autio ends with -us.mp3 or -uk.mp3, then use it
        if (audio.endsWith('-us.mp3') || audio.endsWith('-uk.mp3')) {
            let country = audio.split('-')[1].split('.')[0];
            phonetic_map[country] = {
                'phonetic': phonetic,
                'audio': audio
            }
            continue;
        }
        phonetic_map['default'] = {
            'phonetic': phonetic,
            'audio': audio
        }
    }
    if (phonetic_map.length == 0) {
        return;
    }
    const result = {
        word: word,
        phonetic: json['phonetic'],
        phonetic_map: phonetic_map,
        meanings: json['meanings']
    }
    html = buildDisplayElem(result);
    $('body').css('overflow', 'hidden');
    swal({
        content: html,
    }).then((value) => {
        console.log(value);
        $('body').css('overflow', 'scroll');
    })
}

function buildDisplayElem(result) {
    let content = $('<div style="max-height: 440px">');
    content.append($('<p style="text-align: left; margin-bottom: 5px; font-size: larger; padding-left: 10px">').append($('<span>').append($('<strong>').text(result['word']))));

    let phonetic = $('<p style="text-align: left; margin: 0px">');
    if (result['phonetic_map'].length == 0) {
        phonetic.append($('<span>').text(result['phonetic']));
    } else {
        const country_list = ['uk', 'us', 'default']
        for (let i = 0; i < country_list.length; i++) {
            let country = country_list[i];
            if (result['phonetic_map'][country]) {
                phonetic.append($('<span style="margin-right: 6px; color: dodgerblue;">').text('♫').click(function () {
                    new Audio(result['phonetic_map'][country]['audio']).play()
                }));
                if (country != 'default') {
                    phonetic.append($('<span style="margin-right: 3px; color: darkgray; font-size: small">').text(country));
                }
                phonetic.append($('<span style="margin-right: 20px; color: darkgray; font-size: small">').text(result['phonetic_map'][country]['phonetic']).click(function () {
                    new Audio(result['phonetic_map'][country]['audio']).play()
                }));
            }
        }
    }
    content.append(phonetic)
    content.append($('<hr style="height: 2px; border: 0; background-color: #e8e8e8"/>'))
    let meanings = result['meanings'];
    let meaning_p = $('<div style="max-height: 400px; overflow: scroll">')
    for (let i = 0; i < meanings.length; i++) {
        let meaning = meanings[i];
        if (i > 0) {
            meaning_p.append($('<hr style="text-align: left;margin: 10 0 10 0;width: 90%;background-color: #f2f2f2;border: 0;height: 2px;" />'))
        }
        let part = $('<p style="text-align: left; margin: 0px">');
        part.append($('<span style="color: brown">').text(meaning['partOfSpeech']));
        let def_p = $('<table>')
        let defs = meaning['definitions'];
        for (let j = 0; j < defs.length; j++) {
            let def = defs[j];
            let def_tr = $('<tr>')
            def_tr.append($('<td style="vertical-align: top">').text((j + 1) + '.'))
            def_tr.append($('<td>').text(def['definition']))
            def_p.append(def_tr)
        }
        part.append(def_p);
        meaning_p.append(part)
        def_p.on('mouseup', clickHandler);
    }
    content.append(meaning_p);
    return content[0];
}

$("#main-book").on("mouseup", clickHandler)
