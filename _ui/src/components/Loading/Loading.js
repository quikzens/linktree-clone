import './Loading.css'

export default function Loading(props) {
  const { size, color } = props

  return (
    <div className={`loading ${size} clr-${color}`}>
      <div></div>
      <div></div>
      <div></div>
      <div></div>
    </div>
  )
}
