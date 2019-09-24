import DeviceSelector from './device-selector/container'
import ExplainDevice from './explain-device/container'
import Error from './error/container'
import PaperKey from './paper-key/container'
import PromptReset from './prompt-reset/container'

export const newRoutes = {
  recoverPasswordDeviceSelector: {
    getScreen: (): typeof DeviceSelector => require('./device-selector/container').default,
  },
  recoverPasswordError: {
    getScreen: (): typeof Error => require('./error/container').default,
  },
  recoverPasswordExplainDevice: {
    getScreen: (): typeof ExplainDevice => require('./explain-device/container').default,
  },
  recoverPasswordPaperKey: {
    getScreen: (): typeof PaperKey => require('./paper-key/container').default,
  },
  recoverPasswordPromptReset: {
    getScreen: (): typeof PromptReset => require('./prompt-reset/container').default,
  },
}
