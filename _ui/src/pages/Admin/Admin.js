import LinkEditor from './LinkEditor/LinkEditor'

export default function Admin() {
  return (
    <div className="columns">
      <div className="column p-6">
        <LinkEditor />
      </div>
      <div className="column is-one-third p-6">
        <div className="link-preview"></div>
      </div>
    </div>
  )
}
