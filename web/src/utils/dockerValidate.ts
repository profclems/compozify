export interface DockerValidateResult {
  error: boolean
  message: string
  commad?: string
}

/**
 * dockerValidate - Validate a docker run command
 * @param command - The docker run command to validate
 * @returns DockerValidateResult
 */
export default function dockerValidate(command: string) {
  const dockerRunCommand = command.split(/\s+/g)

  // if less than 2 words, then it's not a valid docker run command
  if (dockerRunCommand.length < 2) {
    return {
      error: true,
      message: 'Invalid command: length need to be 2 or more'
    }
  }

  // if the first word is not `docker`, then it's not a valid docker run command
  if (dockerRunCommand[0] !== 'docker') {
    return {
      error: true,
      message: 'Invalid command: first word need to be `docker`'
    }
  }

  return {
    error: false,
    message: 'Valid docker run command',
    command: dockerRunCommand.join(' ')
  }
}
