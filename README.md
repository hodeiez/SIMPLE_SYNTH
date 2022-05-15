# Simple Synth with go 
A simple synth made using golang 

**under development**
<html>
<body>
<h1>development status</h1>
<p>At the moment: </p>
<ul>
<li>a simple oscillator runs and it can change the frequency by midi notes</li>
<li>four wave type can be selected (square wave needs to be fixed)</li>
<li>A simple gui running where wave type selector is functional</li>
<li>A simple ADSR working controlled via GUI</li>
<li>Basic configurable polyphony </li>
</ul>
<image src="synthPic.png"/>
<h5>actual project status (polyphonic, has a wave selector(sine,triangle, saw, square),ADSR envelope, pitch shift)</h5>
<h1>About</h1>
<p>Before running: get dependencies, and connect a midi controller</p>
<p>Dependencies needed in your OS apart from "go get $whatever":</p>
<ul>
<li>rtmidi</li>
<li>portaudio</li>
</ul>
<h2>Future features</h2>
<ul>
<li>Velocity sensitivity</li>
<li>AM/FM modulation</li>
<li>Low pass effect</li>
<li>Delay effect</li>
<li>Volume control</li>
<li>Dual oscillator</li>
<li>LFO</li>
</ul>
</body></html>