import './App.css'
import Editor, { DiffEditor, useMonaco, loader } from '@monaco-editor/react'
import { useEffect, useRef, useState } from 'react'
import { NavbarComponent } from './Navbar'
import { getAllFilesInRepo } from './utils/github'
import Card from 'react-bootstrap/Card'

function App() {
  const editorRef = useRef(null)
  const [executionState, setExecutionState] = useState([])
  const [val, setVal] = useState(`
  package main

import("fmt")

func main()  {
	fmt.Println("hello world")
}
  `)
  const [files, setFiles] = useState([
    {
      fileName: 'null',
      resourceURI: null
    }
  ])
  async function handleEditorDidMount(editor, monaco) {
    const filesList = await getAllFilesInRepo()
    setFiles(filesList)
    editorRef.current = editor
  }
  function handleEditorChange(value, event) {
    setVal(value)
  }
  useEffect(() => {}, [])
  return (
    <>
      <NavbarComponent
        files={files}
        val={val}
        setVal={setVal}
        setExecState={setExecutionState}
        execState={executionState}
      />
      <Editor
        height="70vh"
        defaultLanguage="go"
        value={val}
        theme="hc-black"
        onMount={handleEditorDidMount}
        onChange={handleEditorChange}
      />
      <Card
        body
        bg="dark"
        text="white"
        style={{ height: '20vh', margin: '10px', overflowY: 'scroll' }}
      >
        {executionState.map((v, i) => {
          return <div key={i}>{v}</div>
        })}
      </Card>
    </>
  )
}

export default App
