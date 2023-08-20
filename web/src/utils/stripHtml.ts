/**
 * stripHtml - strips html tags from a string
 * @param {string} html - html string
 * @returns {string} - string without html tags
 */
export function stripHtml(html: string): string {
  let text = html.replace(/(<([^>]+)>)/gi, '')

  // Create a map of special characters to their text equivalent.
  const specialCharacters = {
    '"': '&quot;',
    "'": '&apos;',
    '<': '&lt;',
    '>': '&gt;',
    '&': '&amp;'
  }

  // Replace all special characters in the text with their text equivalent.
  for (const [character, replacement] of Object.entries(specialCharacters)) {
    text = text.replaceAll(replacement, character)
  }

  return text
}
