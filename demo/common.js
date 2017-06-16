  function makeA(text, href) {
    return '<a  href="' + href + '"' + '>' + text + '</a>'
  }

  function makeFunctionA(text, onClick) {
    return '<a  href="#" onClick="' + onClick + '"' + '>' + text + '</a>'
  }

  function getURLParameter(name) {
    return decodeURIComponent((new RegExp('[?|&]' + name + '=' + '([^&;]+?)(&|#|;|$)').exec(location.search) || [null, ''])[1].replace(/\+/g, '%20')) || null;
  }

  function trimToOneRow(longString) {
    if (longString.length > 130) {
      return longString.substring(0, 126) + "..."
    }
    return longString
  }

  function trimTo40(longString) {
    if (longString.length > 40) {
      return longString.substring(0, 36) + "..."
    }
    return longString
  }

  function noParagraphs(str) {
    str = str.replace(new RegExp("<p>", 'g'), "")
    str = str.replace(new RegExp("</p>", 'g'), "")
    str = str.replace(new RegExp("<br>", 'g'), "")
    str = str.replace(new RegExp("<b>", 'g'), "")
    str = str.replace(new RegExp("</b>", 'g'), "")
    str = str.replace(new RegExp("<blockquote>", 'g'), "")
    str = str.replace(new RegExp("</blockquote>", 'g'), "")    
    return str
  }
