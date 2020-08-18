import requireContext from 'require-context.macro';

// https://github.com/diegohaz/arc/wiki/Atomic-Design#do-not-worry
const req = requireContext('.', true, /\.*\/[^/]+\/[^/]+\/index\.jsx$/)

req.keys().forEach((key) => {
  const componentName = key.replace(/^.+\/([^/]+)\/index\.jsx/, '$1')
  module.exports[componentName] = req(key).default
})
