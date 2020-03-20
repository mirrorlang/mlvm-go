const atomsize=50

function drawatom(card,mem,cpu,i,j){
    let textbox = new fabric.Textbox("", {
        left: j*atomsize,
        top: i*atomsize,
        width: atomsize,
        fontSize: 16, // 字体大小
        // fontWeight: 800, // 字体粗细
        // fill: 'red', // 字体颜色
        // fontStyle: 'italic',  // 斜体
        // fontFamily: 'Delicious', // 设置字体
        // stroke: 'green', // 描边颜色
        // strokeWidth: 3, // 描边宽度
        hasControls: false,
        borderColor: "black",
        bor: "black",
        borderColor: 'orange',
        editingBorderColor: 'blue' // 点击文字进入编辑状态时的边框颜色
    });
    let atom=mem[i][j]
    if(atom!=null)
    switch(atom.Type){
        case "func":
            switch (atom.Op){
                case "func":
                    textbox.text=""+atom.Name;
                    textbox.fontSize=14, // 字体大小
                    textbox.fill="Gold";
        
                    let funcrect= new fabric.Rect( {
                        left: j*atomsize,
                        top: i*atomsize,
                        width:atomsize*atom.Size_x,
                        height:atomsize*atom.Size_y,
                        selectable: false,
                        fill: "rgba(0,0,0,0)",
                        opacity: "1",
                        stroke:"Gold",
                        strokeWidth:1
                        });
                    card.add(funcrect);
                    break
                case "call":
                    let func=mem[atom.Func.Y][atom.Func.X]
                    textbox.text=func.Name+"("
                    
                    // let callnext= new fabric.Line([ j*atomsize,i*atomsize, (j+cpu[0].Funcrect.Size_x)*atomsize,(i)*atomsize], {
                    //     strokeWidth: 1, //线宽
                    //     stroke:"Red", //线的颜色
                    //     selectable: false
                    //     });
                    // card.add(callnext);
                    break   
                case "return":
                    textbox.text=atom.Op
                    break           
            }


            break;
        case "rect":
            textbox.text="□";
            textbox.fill="Navy";

            break;
        case "rectdata":
            textbox.text="↴";
            textbox.fill="SlateBlue";
            break;
        case "point":
            if(atom.Isoffset){
                textbox.text="+"
            }else{
                textbox.text="·"
            }
            textbox.text+=("("+atom.X+","+atom.Y+")")
            textbox.fill="Purple";
            textbox.fontSize=12;

            break;
        case "num":
            textbox.text=atom.IntValue+"";
            textbox.fill="skyblue";
            break;
        case "bool":
            textbox.text=atom.Value+"";
            textbox.fill="skyblue";
            break;
        case "string":
            textbox.text=atom.V_string+"";
            textbox.fill="green";
            break;
        case "controlflow":
            textbox.text=atom.Op+"";
            textbox.fill="red";
            switch(atom.op){
                case "go":
                    let right=mem[i][j+1]
                    textbox.text=atom.Op+""
                    let goline= new fabric.Line([ j*atomsize,i*atomsize, (j+right.Offset_x)*atomsize,(i+right.Offset_y)*atomsize], {
                        strokeWidth: 1, //线宽
                        stroke:"Red", //线的颜色
                        selectable: false
                        });
                    card.add(goline);
                    // let r= new fabric.Rect( {
                    //     left: atom.Point_x*atomsize,
                    //     top:  atom.Point_y*atomsize,
                    //     width:atomsize*atom.Size_x,
                    //     height:atomsize*atom.Size_y,
                    //     selectable: false,
                    //     fill: "rgba(0,0,0,0)",
                    //     opacity: "1",
                    //     stroke:"Navy",
                    //     strokeWidth: 3
                    //     });

                    // card.add(r);
                    break
            }

            break
        case "op":            
            switch (atom.Op){
                case "+":
                case "-":
                case "*":
                case "/":
                    textbox.text=atom.Op+""
                    textbox.fontSize=20;
                    break;
                default:
                    let rectnext= new fabric.Line([ j*atomsize,i*atomsize, (j)*atomsize,(i+1)*atomsize], {
                        strokeWidth: 1, //线宽
                        stroke:"Red", //线的颜色
                        selectable: false
                        });
                    card.add(rectnext);
                    textbox.text=atom.Op+"";
            }
           
            textbox.fill="red";
            break;
    }
// 添加文字后，如下图
    card.add(textbox);
}
function drawcpu(card,cpu){
    let cpup= new fabric.Circle({
        left:cpu.Runfunc.Cpu_x*atomsize,
        top:cpu.Runfunc.Cpu_y*atomsize,
        strokeWidth: 1, //线宽
        radius: atomsize/2,
        stroke:"Red", //线的颜色
        fill: "rgba(0,0,0,0)",
        selectable: false
        });
    card.add(cpup);

    let cpufuncrect= new fabric.Rect( {
        left:cpu.Runfunc.Funcbody.X*atomsize,
        top:cpu.Runfunc.Funcbody.Y*atomsize,
        width:atomsize*cpu.Runfunc.Funcbody.Size_x,
        height:atomsize*cpu.Runfunc.Funcbody.Size_y,
        selectable: false,
        fill: "rgba(0,0,0,0)",
        opacity: "3",
        stroke:"Red",
        strokeWidth:3
        });
    card.add(cpufuncrect);

}
function draw(card,mem,cpu){
    for (let i=0;i<cpu.length;i++){
        drawcpu(card,cpu[i]);
    }
    for (let i=0;i<mem.length;i++){
        for (let j=0;j<mem.length;j++){
            drawatom(card,mem,cpu,i,j)
        }
    }
}

function memborder(card,size){
    for (let i=0;i<=size;i++){
        let horizontal= new fabric.Line([ i*atomsize,0, i*atomsize,size*atomsize], {
            strokeWidth: 1, //线宽
            stroke:"Gainsboro ", //线的颜色
            selectable: false
            });
            card.add(horizontal);
        let vertical= new fabric.Line([ 0,i*atomsize, size*atomsize,i*atomsize], {
            strokeWidth: 1, //线宽
            stroke:"Gainsboro ", //线的颜色
            selectable: false
            });
            card.add(vertical);
    }
   
}