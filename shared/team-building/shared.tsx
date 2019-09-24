import * as Styles from '../styles'
import {ServiceIdWithContact, ServiceMap} from '../constants/types/team-building'
import {IconType} from '../common-adapters/icon.constants-gen'
import Flags from '../util/feature-flags'
import {allServices} from '../constants/team-building'

const serviceColors: {[K in ServiceIdWithContact]: string} = {
  get email() {
    return Styles.isDarkMode ? '#3663ea' : '#3663ea'
  },
  get facebook() {
    return Styles.isDarkMode ? '#3B5998' : '#3B5998'
  },
  get github() {
    return Styles.isDarkMode ? '#333' : '#333'
  },
  get hackernews() {
    return Styles.isDarkMode ? '#FF6600' : '#FF6600'
  },
  get keybase() {
    return Styles.isDarkMode ? '#3663ea' : '#3663ea'
  },
  get phone() {
    return Styles.isDarkMode ? '#3663ea' : '#3663ea'
  },
  get reddit() {
    return Styles.isDarkMode ? '#ff4500' : '#ff4500'
  },
  get twitter() {
    return Styles.isDarkMode ? '#1DA1F2' : '#1DA1F2'
  },
}

const services: {
  [K in ServiceIdWithContact]: {
    icon: IconType
    label: string
    longLabel: Array<string>
    searchPlaceholder: string
    wonderland?: boolean
  }
} = {
  email: {
    icon: 'iconfont-mention',
    label: 'Email', // TODO: rethink this for the empty state when we're actually using it
    longLabel: ['An email', 'address'],
    searchPlaceholder: 'email',
    wonderland: true,
  },
  facebook: {
    icon: 'iconfont-identity-facebook',
    label: 'Facebook',
    longLabel: ['A Facebook', 'user'],
    searchPlaceholder: 'Facebook',
  },
  github: {
    icon: 'iconfont-identity-github',
    label: 'GitHub',
    longLabel: ['A GitHub', 'user'],
    searchPlaceholder: 'GitHub',
  },
  hackernews: {
    icon: 'iconfont-identity-hn',
    label: 'Hacker News',
    longLabel: ['A Hacker', 'News user'],
    searchPlaceholder: 'Hacker News',
  },
  keybase: {
    icon: 'iconfont-contact-book',
    label: 'Keybase and contacts',
    longLabel: Styles.isMobile ? ['Keybase &', 'Contacts'] : ['A Keybase', 'user'],
    searchPlaceholder: Styles.isMobile ? 'Keybase & contacts' : 'Keybase',
  },
  phone: {
    icon: 'iconfont-number-pad',
    label: 'Phone',
    longLabel: ['A phone', 'number'],
    searchPlaceholder: 'phone',
    wonderland: true,
  },
  reddit: {
    icon: 'iconfont-identity-reddit',
    label: 'Reddit',
    longLabel: ['A Reddit', 'user'],
    searchPlaceholder: 'Reddit',
  },
  twitter: {
    icon: 'iconfont-identity-twitter',
    label: 'Twitter',
    longLabel: ['A Twitter', 'user'],
    searchPlaceholder: 'Twitter',
  },
}

const serviceIdToAccentColor = (service: ServiceIdWithContact): string => serviceColors[service]
const serviceIdToIconFont = (service: ServiceIdWithContact): IconType => services[service].icon
const serviceIdToLabel = (service: ServiceIdWithContact): string => services[service].label
const serviceIdToLongLabel = (service: ServiceIdWithContact): Array<string> => services[service].longLabel
const serviceIdToSearchPlaceholder = (service: ServiceIdWithContact): string =>
  services[service].searchPlaceholder
const serviceIdToWonderland = (service: ServiceIdWithContact): boolean =>
  Flags.wonderland && services[service].wonderland === true

const serviceMapToArray = (services: ServiceMap) => allServices.filter(x => x !== 'keybase' && x in services)

export {
  serviceIdToIconFont,
  serviceIdToAccentColor,
  serviceIdToLabel,
  serviceIdToLongLabel,
  serviceIdToSearchPlaceholder,
  serviceIdToWonderland,
  serviceMapToArray,
}
