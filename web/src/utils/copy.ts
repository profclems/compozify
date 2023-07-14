/**
 * copyToClipboardWithMeta - Copies the given value to the clipboard
 * @param value - The value to copy to the clipboard
 * @param _meta - Unused
 */
export default async function copyToClipboardWithMeta(
  value: string,
  _meta?: Record<string, unknown>
) {
  navigator.clipboard.writeText(value)
}
