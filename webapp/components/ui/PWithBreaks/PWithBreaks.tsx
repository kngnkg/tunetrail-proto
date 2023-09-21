import * as React from "react"

export interface PWithBreaksProps {
  text: string
}

export const PWithBreaks: React.FC<PWithBreaksProps> = (props) => {
  const { text } = props
  return (
    <p>
      {text.split("\n").map((str, index, array) => (
        <React.Fragment key={index}>
          {str}
          {index !== array.length - 1 && <br />}
        </React.Fragment>
      ))}
    </p>
  )
}
