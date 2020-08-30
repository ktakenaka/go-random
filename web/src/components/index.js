import requireContext from 'require-context.macro';

// https://github.com/diegohaz/arc/wiki/Atomic-Design#do-not-worry
const req = requireContext('.', true, /\.*\/[^/]+\/[^/]+\/index\.js$/)

req.keys().forEach((key) => {
  const componentName = key.replace(/^.+\/([^/]+)\/index\.js/, '$1')
  module.exports[componentName] = req(key).default
})
